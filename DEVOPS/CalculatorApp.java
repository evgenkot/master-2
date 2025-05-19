import java.util.*;
import java.util.regex.*;

public class CalculatorApp {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("Введите выражение (например, 1+1, sqrt(9), sqr(3)):");
        String input = sc.nextLine();
        try {
            ParsedInput pi = InputParser.parse(input);
            Operation op = OperationFactory.create(pi.op);
            if (pi.binary) ((BinaryOperation) op).setOperands(pi.a, pi.b);
            else ((UnaryOperation) op).setOperand(pi.a);
            double result = op.execute();
            // вывод целого без .0
            if (result == (long) result) System.out.println((long) result);
            else System.out.println(result);
        } catch (CalculatorException e) {
            System.out.println(e.getMessage());
        }
        sc.close();
    }

    // --- Исключение калькулятора ---
    public static class CalculatorException extends Exception {
        public CalculatorException(String message) { super(message); }
    }

    // --- Интерфейсы операций ---
    public interface Operation { double execute() throws CalculatorException; }
    public interface BinaryOperation extends Operation { void setOperands(double a, double b); }
    public interface UnaryOperation extends Operation { void setOperand(double a); }

    // --- Реализации операций ---
    public static class Addition implements BinaryOperation {
        private double a, b;
        public void setOperands(double a, double b) { this.a = a; this.b = b; }
        public double execute() { return a + b; }
    }
    public static class Subtraction implements BinaryOperation {
        private double a, b;
        public void setOperands(double a, double b) { this.a = a; this.b = b; }
        public double execute() { return a - b; }
    }
    public static class Multiplication implements BinaryOperation {
        private double a, b;
        public void setOperands(double a, double b) { this.a = a; this.b = b; }
        public double execute() { return a * b; }
    }
    public static class Division implements BinaryOperation {
        private double a, b;
        public void setOperands(double a, double b) { this.a = a; this.b = b; }
        public double execute() throws CalculatorException {
            if (b == 0) throw new CalculatorException("Невозможно выполнить деление на ноль!");
            return a / b;
        }
    }
    public static class SquareRoot implements UnaryOperation {
        private double a;
        public void setOperand(double a) { this.a = a; }
        public double execute() throws CalculatorException {
            if (a < 0) throw new CalculatorException(
                "Невозможно выполнить извлечение квадратного корня из отрицательного числа!");
            return Math.sqrt(a);
        }
    }
    public static class Square implements UnaryOperation {
        private double a;
        public void setOperand(double a) { this.a = a; }
        public double execute() { return a * a; }
    }

    // --- Фабрика операций ---
    public static class OperationFactory {
        private static final Map<String, Operation> prototypes = new HashMap<>();
        static {
            prototypes.put("+", new Addition());
            prototypes.put("-", new Subtraction());
            prototypes.put("*", new Multiplication());
            prototypes.put("/", new Division());
            prototypes.put("sqrt", new SquareRoot());
            prototypes.put("sqr", new Square());
        }
        public static Operation create(String code) throws CalculatorException {
            Operation op = prototypes.get(code);
            if (op == null) throw new CalculatorException("Неизвестная операция: " + code);
            try {
                return op.getClass().getDeclaredConstructor().newInstance();
            } catch (Exception e) {
                throw new CalculatorException("Ошибка создания операции: " + code);
            }
        }
    }

    // --- Парсер ввода ---
    public static class ParsedInput {
        public final String op; public final double a, b; public final boolean binary;
        public ParsedInput(String op, double a, double b) {
            this.op = op; this.a = a; this.b = b; this.binary = true;
        }
        public ParsedInput(String op, double a) {
            this.op = op; this.a = a; this.b = 0; this.binary = false;
        }
    }

    public static class InputParser {
        // поддерживаем отрицательные и дробные числа
        private static final Pattern BIN = Pattern.compile(
            "\\s*([+-]?[0-9]+(?:\\.[0-9]+)?)\\s*([-+*/])\\s*([+-]?[0-9]+(?:\\.[0-9]+)?)\\s*");
        private static final Pattern UN = Pattern.compile(
            "\\s*(sqrt|sqr)\\s*\\(\\s*([+-]?[0-9]+(?:\\.[0-9]+)?)\\s*\\)\\s*");
        public static ParsedInput parse(String input) throws CalculatorException {
            Matcher m = BIN.matcher(input);
            if (m.matches()) {
                double a = Double.parseDouble(m.group(1));
                String op = m.group(2);
                double b = Double.parseDouble(m.group(3));
                return new ParsedInput(op, a, b);
            }
            m = UN.matcher(input);
            if (m.matches()) {
                String op = m.group(1);
                double a = Double.parseDouble(m.group(2));
                return new ParsedInput(op, a);
            }
            throw new CalculatorException("Ошибка ввода!");
        }
    }
}
