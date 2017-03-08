#!/usr/bin/env python

banner = """
                                     __  __    __                     
                                    /  |/  |  /  |                    
  _______   ______    ______    ____$$ |$$/  _$$ |_                   
 /       | /      \  /      \  /    $$ |/  |/ $$   |                  
/$$$$$$$/ /$$$$$$  |/$$$$$$  |/$$$$$$$ |$$ |$$$$$$/                   
$$ |      $$ |  $$/ $$    $$ |$$ |  $$ |$$ |  $$ | __                 
$$ \_____ $$ |      $$$$$$$$/ $$ \__$$ |$$ |  $$ |/  |                
$$       |$$ |      $$       |$$    $$ |$$ |  $$  $$/                 
 $$$$$$$/ $$/        $$$$$$$/  $$$$$$$/ $$/    $$$$/                  
                                                                      
                                                                      
                                                                      
                                     __                               
                                    /  |                              
  _______   ______    ______    ____$$ |                              
 /       | /      \  /      \  /    $$ |                              
/$$$$$$$/  $$$$$$  |/$$$$$$  |/$$$$$$$ |                              
$$ |       /    $$ |$$ |  $$/ $$ |  $$ |                              
$$ \_____ /$$$$$$$ |$$ |      $$ \__$$ |                              
$$       |$$    $$ |$$ |      $$    $$ |                              
 $$$$$$$/  $$$$$$$/ $$/        $$$$$$$/                               
                                                                      
                                                                      
                                                                      
           __                            __                           
          /  |                          /  |                          
  _______ $$ |____    ______    _______ $$ |   __   ______    ______  
 /       |$$      \  /      \  /       |$$ |  /  | /      \  /      \ 
/$$$$$$$/ $$$$$$$  |/$$$$$$  |/$$$$$$$/ $$ |_/$$/ /$$$$$$  |/$$$$$$  |
$$ |      $$ |  $$ |$$    $$ |$$ |      $$   $$<  $$    $$ |$$ |  $$/ 
$$ \_____ $$ |  $$ |$$$$$$$$/ $$ \_____ $$$$$$  \ $$$$$$$$/ $$ |      
$$       |$$ |  $$ |$$       |$$       |$$ | $$  |$$       |$$ |      
 $$$$$$$/ $$/   $$/  $$$$$$$/  $$$$$$$/ $$/   $$/  $$$$$$$/ $$/       
                                                                      
                                                                      
                                                                      
Welcome to the Credit Card Checker.

Using advanced maximum technology, your credit card will be checked! 

Please enter in your cerdit card to check if it's safe!


"""
print(banner)

name = input("Enter your name: ")
cc = input("Enter your 16 digit credit card number: ")
expiry = input("Enter the expiry date (mm/yy): ")
cvc = input("Enter the CVC on the back: ")

yesno = input("Are you sure to check? ")

print("Thank you. hee hee")
