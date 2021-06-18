## Mocking in GO
Mocking is a process used in unit testing when the unit being tested has external dependencies. 
The purpose of mocking is to isolate and focus on the code being tested and not on the behavior or state of external dependencies.
Mocking is creating objects that simulate the behavior of real objects.

Example: If we instead want to verify that the object under test writes some specific data to the database we will have to mock the database. 

### NOTE: Now Since we are building “Our Service” we want to write an independent unit test for “Our Service” but as we use the functionality of third-party service or library in our service if we test without mock’s we will be doing the integration testing which is sometimes complex and more time-consuming.

About Mocking with Example: [How to Mocking](!https://medium.com/@ankur_anand/how-to-mock-in-your-go-golang-tests-b9eee7d7c266)
1.Mock objects meet the interface requirements:
 First, we have to define our interface requirement that our mock going to implement.  
(need an interface that is just internal to the package and provides a way to test if a user exists or not.)

```
    // registrationPreChecker validates if user is allowed to register.
    type registrationPreChecker interface {
        userExists(string) bool
    } 
```
Interface Implementation:userExists function of our new defined interface simply 
wraps the call to the actual call to third party service.
