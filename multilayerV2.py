import argparse



def can_open(filename):
    try:
        with open(filename, "r") as buff:
            lines = buff.read()
    except FileNotFoundError:
        raise FileNotFoundError
    except PermissionError:
        raise PermissionError
    except IsADirectoryError:
        raise IsADirectoryError
    return lines


def get_args():
    """
    Parse the input
    """
    parser = argparse.ArgumentParser()
    parser.add_argument("data", type=str, help="The data file")
    parser.add_argument("structure", type=str, help="The structure file")
    parser.add_argument("--save", type=str,
                        help="Name of the saving file for weigths", default="")
    parser.add_argument("--load", type=str,
                        help="Name of the loading file for weigths",
                        default="")
    parser.add_argument("--iteration", "-i", type=int,
                        help="Number of iterations that need to be done",
                        default=1000)
    parser.add_argument("--regularisation", "-r", type=float,
                        help="The value of the regularisation in the cost",
                        default=0.016)
    parser.add_argument("--alpha", "-a", type=float,
                        help="the value of alpha. By default,\
                                alpha is dynamic",
                        default=0)
    return parser.parse_args()


def main():
    """
    Main function
    """
    args = get_args()
    data = can_open(args.data)


if __name__ == "__main__":
    main()
