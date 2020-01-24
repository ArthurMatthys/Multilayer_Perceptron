import argparse
import numpy as np


def get_args():
    """
    Parse the input
    """
    parser = argparse.ArgumentParser()
    parser.add_argument("input", type=str,
                        help="The name of the csv file")
    parser.add_argument("--test", "-t",
                        help="file for testing ?", action="store_true")

    return parser.parse_args()


def load_csv(filename):
    """
    Read the csv file
    """
    result = []
    with open(filename) as csv_file:
        lines = csv_file.readlines()
        for line in lines:
            line = line.strip()
            result.append(line.split(','))
    return result


def randomize(a, b):
    """
    Generate the same random permutation on two np array
    """
    permutation = np.random.permutation(a.shape[0])
    shuffled_a = a[permutation]
    shuffled_b = b[permutation]
    return shuffled_a, shuffled_b


def create_npy(data, test):
    """
    Create the input and output files required by the NN to run
    """
    nbr_line, nbr_col = len(data), len(data[0])
    output_array = []
    input_array = []
    for i in range(nbr_line):
        new_line = []
        for j in range(1, nbr_col):
            if j == 1:
                if data[i][j] == 'M':
                    output_array.append([1, 0])
                else:
                    output_array.append([0, 1])
            else:
                new_line.append(float(data[i][j]))
        input_array.append(new_line)
    input_array = np.array(input_array)
    output_array = np.array(output_array)
    max_columns = np.amax(input_array, axis=0)

    for index, max_col in enumerate(max_columns):
        input_array[:, index] /= max_col

    if test:
        np.save("Input_test.npy", input_array)
        np.save("Output_test.npy", output_array)
        return

    split = round(nbr_line * 0.8)
    print(output_array)
    output_array, input_array = randomize(output_array, input_array)
    np.save("Input_training.npy", input_array[:split])
    np.save("Input_training_test.npy", input_array[split:])
    np.save("output_training.npy", output_array[:split])
    np.save("output_training_test.npy", output_array[split:])

if __name__ == "__main__":
    args = get_args()
    create_npy(load_csv(args.input), args.test)
