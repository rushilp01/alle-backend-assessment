# Text-Streaming-Endpoint
----------------------------------------------------
This is a text streaming endpoint. I have already stubbed datasets for 3 providers that we will switch based on a failure condition.

## To run the program:
1. Clone repo
2. install Golang
3. for windows: choco install make (install choco from here https://chocolatey.org/install)
4. run 
```make gorun``` command
5. go to your browser and simply type 
```localhost:8080/stream```.

## Assumptions
1. Used just 3 providers/inferences
2. Have a predefined dataset to choose text from
3. Switching between the 2 based on a failure condition(provider in which the failure rate is the least, will be the new provider).


## Example screenshot of real time provider switching
![Screenshot (4)](https://github.com/user-attachments/assets/4048eaae-b180-4cb1-af40-86c1940d00d1)
