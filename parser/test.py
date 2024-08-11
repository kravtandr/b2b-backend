# read text file line by line
def read_file_line_by_line(file_path):
    count = 0 
    # generate list from 1 to 1001
    list = [x for x in range(0, 1001)]
    with open(file_path, 'r') as file:
        for line in file:
            # Do something with each line here
            if int(line.strip()) in list:
                print(int(line.strip()))
                list.remove(int(line.strip()))
                count+=1
    print(list, sep='\n', end='\n')

def main ():
    read_file_line_by_line('tt.txt')


if __name__ == "__main__":
     main()
