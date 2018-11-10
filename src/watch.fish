for i in (seq 1000)
    echo (set_color blue) (date) waiting... (set_color normal)
    inotifywait -e modify ./*.go
    echo (set_color blue) (date) Compiling....
    gopherjs build -m ./*.go -o ../app/script.js
    echo (set_color green) (date) Completed! (set_color normal)
    echo 
    echo $i
    echo "-----------"
end
