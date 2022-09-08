#include <stdio.h>

#define MAX_TASKS 50000

void process() {
	// input
	int n;
	scanf("%d", &n);
	int a[n];
	for(int i=0;i<n;i++)
		scanf("%d", &a[i]);

	// create array of tasks:
	// if task[i] == 0 that means task #i was not processed before
	// if it was being processed then we are to set task[i] = 1 
	int task[MAX_TASKS] = {0};

	// test sequence
	for(int i=0;i<n;i++) {
		int task_num = a[i];

		// if task was being processed and previous task is not the same task
		if(task[ task_num ] == 1 && i>0 && a[i-1] != task_num) {
			printf("NO\n");
			return;
		}
		task [ task_num ] = 1;
	}
	printf("YES\n");
}

int main() {
	int t;
	scanf("%d", &t);
	for(int i=0;i<t;i++) {
		process();		
	}
	return 0;
}