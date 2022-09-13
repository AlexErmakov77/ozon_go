#include <stdio.h>

// Check if value is in list
int value_in_list(int value, int list[], int n) {
	for(int i=0;i<n;i++) 
		if(list[i]==value) 
			return 1;
	return 0;
}

void process() {
	// input
	int n;
	scanf("%d", &n);
	int next1[n+1], next2[n+1];
	int value;
	int value_list[n+1][2];
	for(int i=1;i<=n;i++) {
		scanf("%d", &value);
		value_list[i][0] = value;
		scanf("%d%d", &next1[value], &next2[value]);
	}

	// for (int i=1;i<=n;i++) {
	// 	printf("%d %d\n", value_list[i][0], value_list[i][1]);
	// }
//	printf("readed data\n");
//
//	printf("n = %d\n", n);
//	for(int i=1;i<=n;i++) {
//		printf("%d %d %d\n", i, next1[i], next2[i]);
//	}

	// recreate list from input data
	int list[n+1];
	list[0] = 1;
	int size = 1;

	while(1) {
	
//		printf("list: ");
//		for(int i=0;i<size;i++) {
//			printf("%d ", list[i]);
//		}
//		printf("\n");

		int index = list[size-1];
		int n1 = next1[index];
		int n2 = next2[index];
		
//		printf("next elements: %d and %d\n", n1, n2);

		if(!value_in_list(n1, list, size)) {
			list[size++] = n1;		
		} else if(!value_in_list(n2, list, size)) {
			list[size++] = n2;		
		} else {
			break;
		}
	}

	// printf("List: ");
	// for(int i=0;i<size;i++) {
	// 	printf("%d ", list[i]);
	// }
    // printf("n = %d\n", n);
	for(int i=1;i<=n+1;i++) {
		printf("%d %d\n", value_list[i][0], value_list[i][1]);
		for (int j=0; j<size; j++) {
			if (value_list[i][0] == list[j]) {
				if (j+size/2 >= size) {
					value_list[i][1] = list[j-size/2];
				} else {
					value_list[i][1] = list[j+size/2];
				}

			}
		}


	}
	for(int i=1;i<n+1;i++) {
		printf("%d %d\n", value_list[i][0], value_list[i][1]);
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