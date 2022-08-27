# ascii-art
```
 _______   _                            _                                     _   _                  _       _           _  
|__   __| (_)                          | |                                   (_) | |                | |     (_)         | | 
   | |     _   _ __ ___     ___        | |_    ___         __      __  _ __   _  | |_    ___        | |__    _    __ _  | | 
   | |    | | | '_ ` _ \   / _ \       | __|  / _ \        \ \ /\ / / | '__| | | | __|  / _ \       | '_ \  | |  / _` | | | 
   | |    | | | | | | | | |  __/       \ |_  | (_) |        \ V  V /  | |    | | \ |_  |  __/       | |_) | | | | (_| | |_| 
   |_|    |_| |_| |_| |_|  \___|        \__|  \___/          \_/\_/   |_|    |_|  \__|  \___|       |_.__/  |_|  \__, | (_) 
                                                                                                                  __/ |     
                                                                                                                 |___/      
```                                                                                                                 
## Usage


1. As first argument program takes `string` you would like to change:

```
/ascii-art$ go run . "Hello"
```
2. You may choose banner template for your `string` by adding `flags`:

Standard(-st, standard):

```
/ascii-art$ go run . "Hello" -st
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```                                
Shadow(-sh, shadow):

```
/ascii-art$ go run . "Hello" -sh
                                 
_|    _|          _| _|          
_|    _|   _|_|   _| _|   _|_|   
_|_|_|_| _|_|_|_| _| _| _|    _| 
_|    _| _|       _| _| _|    _| 
_|    _|   _|_|_| _| _|   _|_|   
                                 
                                 
```
Thinkertoy(-th, thinkertoy):

```
/ascii-art$ go run . "Hello" -th
                 
o  o     o o     
|  |     | |     
O--O o-o | | o-o 
|  | |-' | | | | 
o  o o-o o o o-o 
                 
                 
```

* If no `flag` was given program sets by default standard banner template

```
/ascii-art$ go run . "Hello"
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```                              

#### Authors

@DarkhanShakhan

@arturzhamaliyev