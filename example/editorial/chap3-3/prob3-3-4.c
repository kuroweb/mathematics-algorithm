#include <stdio.h>

long long N;
long long A[200009];
long long a = 0, b = 0, c = 0, d = 0; // オーバーフロー回避のため 64 ビット整数を使う

int main() {
	// 入力
	scanf("%lld", &N);
	for (int i = 1; i <= N; i++) scanf("%lld", &A[i]);
	
	// a, b, c, d の個数を数える
	for (int i = 1; i <= N; i++) {
		if (A[i] == 100) a += 1;
		if (A[i] == 200) b += 1;
		if (A[i] == 300) c += 1;
		if (A[i] == 400) d += 1;
	}
	
	// 出力（答えは a * d + b * c）
	printf("%lld\n", a * d + b * c);
	return 0;
}