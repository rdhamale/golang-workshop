class MultithreadingDemo extends Thread 
{ 
    public void run() 
    { 
        try
        { 
            // Displaying the thread that is running 
            System.out.println ("Thread " + 
                  Thread.currentThread().getId() + 
                  " is running"); 
  
        } 
        catch (Exception e) 
        { 
            // Throwing an exception 
            System.out.println ("Exception is caught"); 
        } 
    } 
} 
  
// Main Class 
public class Multithread 
{ 
    public static void main(String[] args) 
    { 
        long startTime = System.nanoTime();

        int n = 1000000; // Number of threads 
        for (int i=0; i<n; i++) 
        { 
            MultithreadingDemo mTdemo = new MultithreadingDemo(); 
            mTdemo.start(); 
        } 

        long endTime   = System.nanoTime();
        long totalTime = endTime - startTime;
        System.out.println("Total time is: "+totalTime/1000000+ "ms");
    } 
} 