def initializing():
    print("Daemon started")
    return loading_data

def loading_data():
    print("Initialization complete")
    user_input = input("Enter 'y' to load data: ")
    if user_input.lower() == 'y':
        return processing_data
    else:
        return finalizing

def processing_data():
    print("Data loaded")
    user_input = input("Enter 'y' to continue processing data: ")
    if user_input.lower() == 'y':
        return loading_data
    else:
        user_input = input("Enter 'y' to finalize: ")
        if user_input.lower() == 'y':
            return finalizing
        else:
            return processing_data

def finalizing():
    print("Daemon stopped")
    return None

def main():
    state = initializing
    while state:
        state = state()

if __name__ == '__main__':
    main()