using System;

public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("Hello World!");
    }

    static int FuelRequired(int mass)
    {
      return mass / 3 - 2;
    }
}
// //#! /usr/bin/env dotnet script
