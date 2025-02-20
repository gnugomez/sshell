#!/bin/bash
# A simple interactive menu

clear
echo "Welcome to the Fake SSH Menu"
echo "--------------------------------"
echo "1) Show current date and time"
echo "2) List files in the current directory"
echo "3) Show system information"
echo "4) Exit"
echo "--------------------------------"

read -p "Choose an option [1-4]: " choice

case $choice in
    1)
        echo "Current date and time: $(date)"
        ;;
    2)
        echo "Files in current directory:"
        ls -lah
        ;;
    3)
        echo "System information:"
        uname -a
        ;;
    4)
        echo "Goodbye!"
        exit 0
        ;;
    *)
        echo "Invalid option!"
        ;;
esac

echo ""
read -p "Press Enter to exit..."

