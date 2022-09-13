#include <stdio.h>

void process() {
    int calendar[13] = {29, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    int d, m, y;
    scanf("%d %d %d", &d, &m, &y);
    if ((d>=1 && d<=31) && (m>=1 && m<=12) && (y>=1950 && y<=2300)) {
        if ((y%4 == 0 && y%100 != 0) || y%400 == 0) {
            m = 0;
        }

        if (d <= calendar[m]) {
            printf("YES\n");
        } else {
            printf("NO\n");
        }
    }    
}

int main() {
	int t;
	scanf("%d", &t);
	for(int i=0;i<t;i++) {
		process();		
	}
	return 0;
}