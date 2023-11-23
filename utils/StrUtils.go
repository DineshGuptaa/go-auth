package utils

// Function to check if an array contains a string
func Contains(array []string, str string) bool {
    for _, s := range array {
      if s == str {
        return true
      }
    }
    return false
  }