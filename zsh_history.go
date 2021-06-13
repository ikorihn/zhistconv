package main

// 131は0x83の10進数表現 zsh_historyの特殊仕様
const (
	metachar = 131
	start    = 130
	end      = 158
	xA0      = 160
)

func parseZshHistory(latin1Byte []byte) string {
	isMarking := false
	var byteBuffer []byte

	for _, codePoint := range latin1Byte {
		if codePoint == metachar {
			isMarking = true
			continue
		}

		if isMarking {
			// 6bit目を反転させるために
			// 0x20をXORする
			invertCodePoint := codePoint ^ 32
			byteBuffer = append(byteBuffer, invertCodePoint)
			isMarking = false
		} else {
			byteBuffer = append(byteBuffer, codePoint)
		}
	}

	return string(byteBuffer)
}

func convertToZshHistory(latin1Byte []byte) string {
	var byteBuffer []byte

	for _, codePoint := range latin1Byte {
		isInverse := false
		// 131は0metacharの10進数表現
		if (start < codePoint && codePoint < end) || codePoint == xA0 {
			isInverse = true
		}

		if isInverse {
			// 6bit目を反転させるために
			// 0x20をXORする
			invertCodePoint := codePoint ^ 32
			byteBuffer = append(byteBuffer, metachar)
			byteBuffer = append(byteBuffer, invertCodePoint)
			isInverse = false
		} else {
			byteBuffer = append(byteBuffer, codePoint)
		}
	}

	return string(byteBuffer)
}
