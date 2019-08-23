package main;

import (
    "fmt";
    "os";
    "bufio";
)

func check_flag(user_input []byte) bool {

    scrambled_flag := []string{"!", "-", "0", "0", "0", "3", "4", "A", "F", "G", "G", "L", "L", "T", "W", "_", "_", "c", "e", "g", "l", "m", "n"};

    flag := string(scrambled_flag[8] + scrambled_flag[11] + scrambled_flag[7] + scrambled_flag[9] + scrambled_flag[1] + scrambled_flag[14] + scrambled_flag[5] + scrambled_flag[20] + scrambled_flag[17] + scrambled_flag[2] + scrambled_flag[21] + scrambled_flag[18] + scrambled_flag[15] + scrambled_flag[13] + scrambled_flag[2] + scrambled_flag[15] + scrambled_flag[9] + scrambled_flag[2] + scrambled_flag[11] + scrambled_flag[6] + scrambled_flag[22] + scrambled_flag[19] + scrambled_flag[0])

    return string(user_input) == flag;
}

func main() {
    fmt.Print("`Go` find the flag and enter it here: ");
    user_input,_,err := bufio.NewReader(os.Stdin).ReadLine();
    if err != nil {
            fmt.Println("Invalid input :/ , ",err);
    }
    if check_flag(user_input) {
        fmt.Println("Congrats!");
    } else {
        fmt.Println("That's not the flag :( ");
    }
}
