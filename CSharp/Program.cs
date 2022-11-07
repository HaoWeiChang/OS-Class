using System.Diagnostics;

public class Matrix
{
    private static int Mi;
    private static int Mj;
    private double[,] MatrixA;
    private double[,] MatrixB;
    public Matrix(int i, int j)
    {
        Mi = i;
        Mj = j;
        MatrixA = new double[Mi, Mj];
        MatrixB = new double[Mj, Mi];
        InitMatrixA(ref MatrixA);
        InitMatrixB(ref MatrixB);
    }
    private void InitMatrixA(ref double[,] MatrixA)
    {
        for (int i = 1; i <= MatrixA.GetLength(0); i++)
        {
            for (int j = 1; j <= MatrixA.GetLength(1); j++)
            {
                MatrixA[i - 1, j - 1] = (6.6 * i) - (3.3 * j);
            }
        }
    }
    private void InitMatrixB(ref double[,] MatrixB)
    {
        for (int i = 1; i <= MatrixB.GetLength(0); i++)
        {
            for (int j = 1; j <= MatrixB.GetLength(1); j++)
            {
                MatrixB[i - 1, j - 1] = 100 + (2.2 * i) - (5.5 * j);
            }
        }
    }
    private double Calculate(int i, int j)
    {
        double res = 0;
        for (int x = 0; x < Mj; x++)
        {
            res += MatrixA[i, x] * MatrixB[x, j];
        }
        return res;
    }
    public void forLoop(ref double[,] mtx)
    {
        for (int i = 0; i < mtx.GetLength(0); i++)
        {
            for (int j = 0; j < mtx.GetLength(1); j++)
            {
                mtx[i, j] = Calculate(i, j);
            }
        }
    }
    public void multiplethread1(double[,] mtx)
    {
        List<Thread> threads = new List<Thread>();
        for (int i = 0; i < mtx.GetLength(0); i++)
        {
            for (int j = 0; j < mtx.GetLength(1) / 50; j++)
            {
                Thread thread = new(() =>
                {
                    int start = j * 50;
                    int end = start + 50;
                    for (int x = start; x < end; x++)
                    {
                        mtx[i, j] = Calculate(i, j);
                    }
                });
                thread.Start();
                threads.Add(thread);
            }
        }
        foreach (Thread t in threads)
            t.Join();
        return mtx;
    }
    public void multiplethread2(ref double[,] mtx)
    {

    }
}

class Program
{
    const int Mi = 50;
    const int Mj = 80;
    static double[,] MatrixC = new double[Mi, Mi];
    static void Main()
    {
        Matrix mtx = new(Mi, Mj);
        Stopwatch stopwatch = new();
        stopwatch.Start();
        mtx.forLoop(ref MatrixC);
        stopwatch.Stop();
        Console.WriteLine("Elapsed Time is {0} ms", stopwatch.ElapsedMilliseconds);
    }
}