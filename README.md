# Ollama-images 

Use Ollama to query an image. 

## Running 

![samples/cartoon.png](samples/cartoon.png)

```bash
$ ollama serve &
$ ollama pull llava 
$ ./build.sh -b 
$ ./bin/describe -image=samples/cartoon.png

15:12:27 INFO Model: llava
15:12:27 INFO Query: What is in this image? Use markdown format for your response.

  The image shows a cute, cartoon-style drawing of an animal that resembles a
  pig. It has a happy expression and is waving with one hand. The background
  is plain white, which makes the character stand out.
```
