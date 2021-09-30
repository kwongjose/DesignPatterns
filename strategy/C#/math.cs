using System;

interface IOperation {
    int calculate(int x, int y);
}

class CalcStrat {
    private IOperation operation; 

    public int run(int x, int y) {
        return this.operation.calculate(x, y);
    }

    public void setStrategy(String strategy) {
        switch (strategy) {
            case "add":
                this.operation = new Addition();
                break;
            case "sub":
                this.operation = new Subtraction();
                break;
            default: 
                Console.WriteLine("Invalid option");
                break;
        }
    }
}