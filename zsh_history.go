package zhistconv

const (
	// zsh_historyの仕様で、各バイトが0x83~0xA2のとき、その前に0x83を入れて6bit目を反転させる
	x83 = 131
	xA2 = 162
	x20 = 32
)

func ParseZshHistory(latin1Byte []byte) string {
	isMarking := false
	var byteBuffer []byte

	for _, codePoint := range latin1Byte {
		if codePoint == x83 {
			isMarking = true
			continue
		}

		if isMarking {
			// 6bit目を反転させるために0x20をXORする
			invertCodePoint := codePoint ^ x20
			byteBuffer = append(byteBuffer, invertCodePoint)
			isMarking = false
		} else {
			byteBuffer = append(byteBuffer, codePoint)
		}
	}

	return string(byteBuffer)
}

func ConvertToZshHistory(latin1Byte []byte) string {
	var byteBuffer []byte

	for _, codePoint := range latin1Byte {
		// 131は0metacharの10進数表現
		if x83 <= codePoint && codePoint <= xA2 {
			// 6bit目を反転させるために0x20をXORする
			invertCodePoint := codePoint ^ x20
			byteBuffer = append(byteBuffer, x83)
			byteBuffer = append(byteBuffer, invertCodePoint)
		} else {
			byteBuffer = append(byteBuffer, codePoint)
		}
	}

	return string(byteBuffer)
}
