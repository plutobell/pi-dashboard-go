# Changelog #

**2021-08-14**

* v1.3.3 : 
  * Added automatic recognition and switching of favicon
  * Adjusted static file path structure
  * Adjusted project structure
  * Updated dependencies

**2021-08-13**

* v1.3.2 : 
  * Added automatic version detection and prompting
* v1.3.1 : 
  * Added csrf protection
  * Adjusted some details
  * Updated dependencies

**2021-08-12**

* v1.3.0 : 
  * Adjusted routing structure
  * Adjusted log formatting
  * Added gzip compression
  * Enhanced web security
  * Login request changed to asynchronous
  * Refactored part of the code
  * Updated dependencies

**2021-08-10**

* v1.2.1 : 
  * Added  login empty field checks
  * Added login failure prompt message
  * Added new command line parameter: log

* v1.2.0 : 
  * Rewrite login authentication
  * Added new login page
  * Added new command line parameter: session
  * Use Go v1.16.7
  * Updated dependencies

**2021-06-17**

* v1.1.2 : 
  * Changed the way cpu usage is calculated
  * Updated dependencies

**2021-06-16**

* v1.1.1 : 
  * Fix the bug of abnormal display of cpu information
  * Use Go v1.16.5
  * Updated dependencies

**2021-04-05**

* v1.1.0 : 
  * Replace go.rice with go embed
  * Fix the bug of program error when unable to get cpu model
  * Added new command line parameter: interval
  * Added a different header image for non-Raspberry Pi devices

**2021-03-31**

* v1.0.10 :
  * Fix the bug of panic caused by empty device model information #1
  * Fix the bug of not finding dashboard.min.js #2
  * Fix the bug of invalid hostname command on arch
  * Added golang version display on view

**2020-8-14**

* v1.0.9 : 
  * Fix swap display bug
  * Adapt to linux system under 386 and amd64
  * Optimize error handling for function Popen

**2020-8-9**

* v1.0.8 : 
  * Optimize swap display details
  * Added shortcut buttons such as shutdown and reboot

**2020-8-7**

* v1.0.7 : 
  * Optimize network card flow and curve display
  * Interface detail adjustment

**2020-8-6**

* v1.0.6 : 
  * Fix the bug that the network card data display error
  * Fixed navigation bar at the top
  * Interface detail adjustment
* v1.0.5 : 
  * Interface color adjustment
  * Data update detection and prompt
  * Optimize code for server
  * Detail adjustment
* v1.0.4 : 
  * Adjust Cached calculation method
  * Added theme-color for mobile browser
  * Added display login user statistics
  * Bug fixes and details optimization

**2020-8-5**

* v1.0.3 : 
  * Newly added time formatting function resolveTime
  * Detail optimization
* v1.0.2 : 
  * Improve command line parameter verification
  * Detail optimization
  * Added test case device_test.go
  * New page loading animation

**2020-8-4**

* v1.0.1 : Bug fixes, detail optimization
* v1.0.0