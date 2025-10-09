package utils

func Contains(str, substr string) bool {
	return len(str) >= len(substr) && (str == substr || len(str) > len(substr) && (str[:len(substr)] == substr))
}
