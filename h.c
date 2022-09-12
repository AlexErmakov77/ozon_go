#include <stdio.h>

int process() {

    int n;
    scanf("%d", &n);

    int a[n][4];
   // int d, m, y, res;
    for (int i=0; i<n; i++) {
        scanf("%d %d %d", &a[i][0], &a[i][1], &a[i][2]);
    }
   
    printf("YES\n");

    int next, temp, count=0;

    for (int i=0; i<n; i++) {
        temp = a[i][0];
        printf("temp %d\n", temp);
        count = 0;
        for (int j=0; j<n; j++) {            
            if ((temp == a[j][1]) && (temp != a[j][0])) {
            count++;
            next =  a[j][0];
            temp = a[j][0];    
            } else if  ((temp == a[j][2]) && (temp != a[j][0])) {
            count++;
            next =  a[j][0]; 
            temp = a[j][0];
            }
            printf("* temp %d\n", temp);
            if (count == n/2-1) {
                a[i][3] = next;
                printf("next %d\n", next);
                count = 0;
            }
        }
    }


    for (int i=0; i<n; i++) {
        printf("%d %d\n", a[i][0], a[i][3]);
    }



}

int main() {
	int t;
	scanf("%d", &t);
	for(int i=0;i<t;i++) {
		process();		
	}
    printf("*** YES\n");
	return 0;
}