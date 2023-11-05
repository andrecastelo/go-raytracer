package clamp

type number interface {
    int | float64
}

func Clamp[T number](value T, max_value T, min_value T) T {
    if value < min_value {
        return min_value
    } else if value > max_value {
        return max_value
    } else {
        return value
    }
}
