import (
    "strings"
)

func lengthOfLongestSubstring(s string) int {
    // What needs to be done:
    // 1. Iterate through the string
    // 2. Store data as you iterate:
    //      - The substring up to that point (Reset when running into repeated characters)
    //      - The length of longest substring without repeating characters (up to that point)
    //      
    
    
    if len(s) == 0 {
        return 0
    }

    currentNonRepeatingString := ""
    
    currentLongestLengthWihtoutRepeating := 1

    for _, char := range s {
        
        if strings.Contains(currentNonRepeatingString, string(char)) {
            fmt.Printf("Found repeating character!\n")
            
            // record current highest
            if len(currentNonRepeatingString) > currentLongestLengthWihtoutRepeating  {
                currentLongestLengthWihtoutRepeating = len(currentNonRepeatingString)    
            }
            
            indexOfFirstAppearance := strings.IndexRune(currentNonRepeatingString, char)
            
            // cut off at the first appearance
            currentNonRepeatingString =  currentNonRepeatingString[indexOfFirstAppearance+1:]
            currentNonRepeatingString += string(char)
            

        } else {
            // add current char to the buffer
            currentNonRepeatingString += string(char)
        }
        
        
        
        
        
        // fmt.Printf("character %c starts at byte position %d\n", char, pos)
        fmt.Printf("currentString = %s\n", currentNonRepeatingString)
        fmt.Printf("currentLongestLength = %d\n", currentLongestLengthWihtoutRepeating)
        fmt.Println("-------------------\n")
    }
    
    
    
    if len(currentNonRepeatingString) > currentLongestLengthWihtoutRepeating {
        return len(currentNonRepeatingString)
    } else {
        return currentLongestLengthWihtoutRepeating
    }
    
}


