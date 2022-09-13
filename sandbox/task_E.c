#include <stdio.h>

#define MAX_SEC 24 * 60 * 60

int check_time(int h, int m, int s) {
	return  (h >= 0 && h <= 23 && m >= 0 && m <= 59 && s >= 0 && s <= 59);
}




int process() {
	int result = 1;

	// just array for check intersections
	int sec[ MAX_SEC ] = {0};

	// input
	int n;
	scanf("%d", &n);

	for(int i=0;i<n;i++) {
		int h1, m1, s1, h2, m2, s2;
		if (scanf("%d:%d:%d-%d:%d:%d", &h1, &m1, &s1, &h2, &m2, &s2) != 6) {
			result = 0;
		}
		else if(!check_time(h1, m1, s1) || !check_time(h2, m2, s2)) {
			result = 0;
		}
		else {
			int t1 =  3600 * h1 + 60 * m1 + s1;
			int t2 =  3600 * h2 + 60 * m2 + s2;

			if(t2<t1) {
				result = 0;
			}
			else {
				// check for intersections
				for(int j=t1; j<=t2; j++) {
					if(sec[j] == 1) {
						result = 0; // it's intersection
						break;
					}
					else
						sec[j] = 1;
				}
			}
		}
	}
	return result; // it's ok
}

int main() {
	int t = 1;
	scanf("%d", &t);
	for (int i = 0; i < t; i++) {
		if(process())
			printf("YES\n");
		else
			printf("NO\n");
	}
	return 0;
}