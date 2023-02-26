package but

// Returns "x" if b is true and "" otherwise.
func X(b bool) string {
	if b {
		return "x"
	}
	return ""
}

// Returns "X" if b is true and "" otherwise.
func XU(b bool) string {
	if b {
		return "X"
	}
	return ""
}

// Returns "1" if b is true and "0" otherwise.
func OZ(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

// Returns "true" if b is true and "false" otherwise.
func Tf(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// Returns "True" if b is true and "False" otherwise.
func TF(b bool) string {
	if b {
		return "True"
	}
	return "False"
}

// Returns "TRUE" if b is true and "FALSE" otherwise.
func TFU(b bool) string {
	if b {
		return "TRUE"
	}
	return "FALSE"
}

// Returns "y" if b is true and "n" otherwise.
func Y(b bool) string {
	if b {
		return "y"
	}
	return "n"
}

// Returns "Y" if b is true and "N" otherwise.
func YU(b bool) string {
	if b {
		return "Y"
	}
	return "N"
}

// Returns "yes" if b is true and "no" otherwise.
func Yn(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// Returns "Yes" if b is true and "No" otherwise.
func YN(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

// Returns "YES" if b is true and "NO" otherwise.
func YNU(b bool) string {
	if b {
		return "YES"
	}
	return "NO"
}

// Returns "on" if b is true and "off" otherwise.
func Oo(b bool) string {
	if b {
		return "on"
	}
	return "off"
}

// Returns "On" if b is true and "Off" otherwise.
func OO(b bool) string {
	if b {
		return "On"
	}
	return "Off"
}

// Returns "ON" if b is true and "OFF" otherwise.
func OOU(b bool) string {
	if b {
		return "ON"
	}
	return "OFF"
}
