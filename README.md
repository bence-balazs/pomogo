# pomogo
Dead simple pomodo timer in go. \
Can take 2 arg(s) from cli: \
    ```--time 20``` (time in minutes)\
    ```--tts hello``` (text to speach at the end of the current pomodo session)

## usage
build:
```sh
make build
```
run:
```sh
./pomogo.app --time 20 --tts hello
```

## required packages in linux:
```sh
spd-say
zenity
```
(should be default on every major distro)
