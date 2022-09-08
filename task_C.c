#include <stdio.h>

int abs(int x) {
	return x>0 ? x : -x;
}

int process() {
	// input
	int n;
	scanf("%d", &n);
	int a[n];
	for(int i=0;i<n;i++)
		scanf("%d", &a[i]);
	
	int pairs = n/2;
	for(int i=0;i<pairs; i++) {
		// get first programmer
		int j1 = 0;
		while(a[j1] == 0) ++j1;

		// found a best pair for him
		int min_dif = 100;
		int j2;
		for(int k = j1+1; k<n; k++) {
			if(a[k]!=0 && abs(a[j1]-a[k])<min_dif) {
				j2 = k;
				min_dif = abs(a[j1]-a[k]);
			}
		}

		// print current pair indexes (j1, j2), 
		// plus one because we are to count from 1, not 0
		printf("%d %d\n", j1+1, j2+1);

		// instead of remove them from list we just set their level to 0
		a[j1] = 0;
		a[j2] = 0;
	}
	printf("\n");
}

int main() {
	int t;
	scanf("%d", &t);
	for(int i=0;i<t;i++) {
		process();		
	}
	return 0;
}