# helloworld

This application has been developed with GoLang in which the application is an HTTP server that sends these plain-text responses on these endpoints:
GET / should respond "hello world"       (eg: http://x.x.x.x:8080/)
GET /hello should respond "hello"          (eg: http://x.x.x.x:8080/hello)
GET /world should respond "world"       (eg: http://x.x.x.x:8080/world)
 
Each endpoint should also accept these query parameters:
uppercase, which if true should return the string but capitalized, e.g. HELLO WORLD

{
eg:  http://x.x.x.x:8080/Hello%20World?uppercase=true

Content in the browser will be: 
Hello World Value After Capitilization: HELLO WORLD
}

reversed, which if true should return the reverse of the string, e.g. dlrow olleh

{
eg: http://x.x.x.x:8080/Hello%20World?reverse=true

Content in the browser will be: 
Hello World Value After Reversing: dlroW olleH 
}

{
eg:  http://x.x.x.x:8080/Hello%20World?uppercase=true&reverse=true

Content in the browser will be: 
Hello World Value After Capitilization: HELLO WORLD Value After Reversing: dlroW olleH 
}
