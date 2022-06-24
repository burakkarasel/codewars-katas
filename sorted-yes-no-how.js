function isSortedAndHow(array) {
  return array.join("") === [...array].sort((a, b) => a - b).join("")
    ? "yes, ascending"
    : array.join("") === [...array].sort((a, b) => b - a).join("")
    ? "yes, descending"
    : "no";
}
