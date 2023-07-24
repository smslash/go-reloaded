package goreloaded

import (
	"fmt"
	"strconv"
)

func Modification(sample string) string {
	var result []string
	template := Fields(sample)
	if len(template) == 0 {
		return " "
	}
	boolean := make([]bool, len(template))

	for i := 0; i < len(template); i++ {
		if template[i] == "(hex)" || template[i] == "(bin)" || template[i] == "(up)" || template[i] == "(low)" || template[i] == "(cap)" {
			boolean[i] = true
		} else if template[i] == "(up," || template[i] == "(low," || template[i] == "(cap," {
			ans := true
			if template[i+1][len(template[i+1])-1] == ')' {
				for j := 0; j < len(template[i+1])-1; j++ {
					if template[i+1][j] != '-' && !IsNumber(template[i+1][j]) {
						ans = false
						break
					}
				}
			}
			if ans {
				boolean[i] = true
				boolean[i+1] = true
				i++
			}
		}
	}

	for i := 0; i < len(template); i++ {
		if boolean[i] {
			if template[i] == "(hex)" {
				count := 0
				for k := i - 1; k >= 0; k-- {
					if template[k] == "\n" {
						count++
					}
					if !boolean[k] && template[k] != "\n" {
						result = result[:len(result)-count-1]
						result = append(result, Hex(template[k]))
						for count > 0 {
							result = append(result, "\n")
							count--
						}
						break
					}
				}
			} else if template[i] == "(bin)" {
				count := 0
				for k := i - 1; k >= 0; k-- {
					if template[k] == "\n" {
						count++
					}
					if !boolean[k] && template[k] != "\n" {
						result = result[:len(result)-count-1]
						result = append(result, Bin(template[k]))
						for count > 0 {
							result = append(result, "\n")
							count--
						}
						break
					}
				}
			} else if template[i] == "(up)" {
				count := 0
				for k := i - 1; k >= 0; k-- {
					if template[k] == "\n" {
						count++
					}
					if !boolean[k] && template[k] != "\n" {
						result = result[:len(result)-count-1]
						result = append(result, Up(template[k]))
						for count > 0 {
							result = append(result, "\n")
							count--
						}
						break
					}
				}
			} else if template[i] == "(low)" {
				count := 0
				for k := i - 1; k >= 0; k-- {
					if template[k] == "\n" {
						count++
					}
					if !boolean[k] && template[k] != "\n" {
						result = result[:len(result)-count-1]
						result = append(result, Low(template[k]))
						for count > 0 {
							result = append(result, "\n")
							count--
						}
						break
					}
				}
			} else if template[i] == "(cap)" {
				count := 0
				for k := i - 1; k >= 0; k-- {
					if template[k] == "\n" {
						count++
					}
					if !boolean[k] && template[k] != "\n" {
						result = result[:len(result)-count-1]
						result = append(result, Cap(template[k]))
						for count > 0 {
							result = append(result, "\n")
							count--
						}
						break
					}
				}
			} else if template[i] == "(up," || template[i] == "(low," || template[i] == "(cap," {
				tmp := template[i+1][:len(template[i+1])-1]
				number, err := strconv.Atoi(tmp)
				if err != nil {
					fmt.Println("Error converting string to integer")
					return " "
				} else {
					if number < 0 {
						fmt.Println("Error: the first file has a negative number")
						return " "
					} else if number > 0 {
						word_counter := 0
						for v := 0; v < i; v++ {
							if template[v] != "\n" && !boolean[v] {
								word_counter++
							}
						}
						if word_counter < number {
							fmt.Println("Error: in the first file the specified number is too large")
							return " "
						}
						if template[i] == "(up," {
							var result1 []string
							for k := 0; k < len(result); k++ {
								if result[k] != "\n" {
									result1 = append(result1, result[k])
								}
							}
							result1_count := 1
							for k := len(result) - 1; k >= 0; k-- {
								if result[k] != "\n" {
									result[k] = Up(result1[len(result1)-result1_count])
									result1_count++
									number--
									if number == 0 {
										break
									}
								}
							}
						} else if template[i] == "(low," {
							var result1 []string
							for k := 0; k < len(result); k++ {
								if result[k] != "\n" {
									result1 = append(result1, result[k])
								}
							}
							result1_count := 1
							for k := len(result) - 1; k >= 0; k-- {
								if result[k] != "\n" {
									result[k] = Low(result1[len(result1)-result1_count])
									result1_count++
									number--
									if number == 0 {
										break
									}
								}
							}
						} else if template[i] == "(cap," {
							var result1 []string
							for k := 0; k < len(result); k++ {
								if result[k] != "\n" {
									result1 = append(result1, result[k])
								}
							}
							result1_count := 1
							for k := len(result) - 1; k >= 0; k-- {
								if result[k] != "\n" {
									result[k] = Cap(result1[len(result1)-result1_count])
									result1_count++
									number--
									if number == 0 {
										break
									}
								}
							}
						}
					}
				}
			}
		} else {
			result = append(result, template[i])
		}
	}

	var final_res string
	for _, value := range result {
		final_res += value + " "
	}
	final_res = final_res[:len(final_res)-1]

	for i := 0; i < len(final_res); i++ {
		if i+1 < len(final_res) && final_res[i+1] == '\n' {
			tmp1 := final_res[:i]
			tmp2 := final_res[i+1:]
			final_res = tmp1 + tmp2
		}
		if i+2 < len(final_res) && final_res[i] == '\n' && final_res[i+1] == ' ' {
			tmp1 := final_res[:i+1]
			tmp2 := final_res[i+2:]
			final_res = tmp1 + tmp2
		}
	}

	return final_res
}
