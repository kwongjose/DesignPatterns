class Main {
    static void Main(string[] args) {
        CalcStrat proxy = new CalcStrat();

        proxy.setStrategy("add");
        int result = proxy.run(1, 2);
        Console.WriteLine($"1 + 2 = {result}");

        proxy.setStrategy("sub");
        result = proxy.run(result, 2);
        Console.WriteLine($"3 - 2 = {result}");
    }
}