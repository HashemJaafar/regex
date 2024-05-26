import os
import re
from collections import Counter

# Set the directory you want to search
directory = r"C:\Users\hashem\Desktop\programing\my_big_project\dht"  # Adjust this path as needed

# Function to read all files in the directory
def read_files_in_directory(directory):
    words = []
    for root, dirs, files in os.walk(directory):
        for file in files:
            if file.endswith(".go"):  # Adjust the file extension as needed
                with open(os.path.join(root, file), 'r', encoding='utf-8', errors='ignore') as f:
                    words.extend(re.findall(r'\w+', f.read()))
    return words

# Get words from files
words = read_files_in_directory(directory)

# Count word occurrences
word_count = Counter(words)

# Sort and display the words by frequency
# sorted_words = word_count.most_common()
sorted_words = sorted(word_count.items())


#####################################################

# # print the words and their counts
# for word, count in sorted_words:
#     print(f"{count}: {word}")

# Write the words and their counts to a text file
output_file = os.path.join(directory, 'word_count.txt')
with open(output_file, 'w', encoding='utf-8') as f:
    for word, count in sorted_words:
        f.write(f"{count}: {word}\n")

print(f"Word counts have been written to {output_file}")