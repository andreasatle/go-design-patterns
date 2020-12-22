# Singleton Design Pattern

Singleton Design Pattern is a creational design pattern and also one of the most commonly used design pattern. This pattern is used when only a single instance of the struct should exist. This single instance is called a singleton object. Some of the cases where the singleton object is applicable:

1. DB instance – we only want to create only one instance of DB object and that instance will be used throughout the application. 
2. Logger instance – again only one instance of the logger should be created and it should be used throughout the application.

The singleton instance is created when the struct is first initialized.  Usually, there is getInstance() method defined on the struct for which only one instance needs to be created. Once created then the same singleton instance is returned every time by the GetInstance().

## Singleton Design Pattern using sync.Mutex
We need to be very careful and check if the singleton is created both before and after using the Lock.
## Singleton Design Pattern using init function
When the singleton can be created at the start of the program, then
we can use the init-function for the creation.
## Singleton Design Pattern using sync.Once
When we want to delay the creation to when the singleton is first used, then we can use sync.Once.