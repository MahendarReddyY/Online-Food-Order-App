This is a simple command Line interface application for placing a food order in a whatAburger

I have used Go lang to develop this one.

This project contains around 10 packages including contact, location, menulist, newyork, payment, promotional code,
receipt generation file.

I have implemented Go routines in this project for running multi Threads and Concurrency.

##The workflow I was created is :

1. When programm starts, it will prompt the user to select the location or Area code to select a  restaurnt , \navigated\
I have insert Newyork and california as two 
locations so far,

2.When user sucessfully enter the city, then the user navigated to the Homepage, 
if invalid entry of the city, it will prompt the user to re enter the location.

3.In Home page it will display 3 items the one is todays menu , contact, offers.

4. when user clicks on the menu it will display the foods and user will add them into the cart and checkout

5. when comes to the payment, After a sucessfully payment it will display the order summary and a recipt(not for Print)

In new version:
1) I have added a promocode feature, when you apply a promocode it will match the code if correct match happens,
  -   Successfully applied the promocode and you can see the changes in your total price 
  -   If no match occured displays the previous prices only 

2) At the end I printed the receipt it will get the details of the summary of your order

Database Connectivity:

I have use the Go database drivers like JDBC connection in java to connect the local database to store the data and used for the further operations.
