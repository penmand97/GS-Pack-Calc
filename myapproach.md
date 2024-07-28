# Pack Consolidation Web Application

Welcome to the my submission for the Gymshark Pack Consolidation Assesment!

This guide should help you understand my thought process, the methods I used, and my approach to creating this project. I will explain everything in simple terms, so it's easy to follow along.

## Thought Process

1. **Understanding the Problem**:
    - I wanted to create a tool that helps us pack items efficiently using different pack sizes
    - The tool should also display the pack sizes we have and the final consolidated packs without showing packs with a value of zero.

2. **Breaking Down the Task**:
    - First, I need to take input from the user for the pack sizes and the number of items.
    - Then, I need to process this input to calculate the best way to pack the items.
    - Finally, I will display the results on the web page.

3. **User Interface Design**:
    - My web page will have two forms: one for entering pack sizes and another for entering the number of items.
    - I will display the pack size options and the consolidated packs as results on the same page.
    - I also want to include a responsive logo at the top of the page.

## Methods Used

1. **HTML (HyperText Markup Language)**:
    - I use HTML to create the structure of my web page.
    - It defines elements like headings, forms, input fields, and buttons.

2. **CSS (Cascading Style Sheets)**:
    - I use CSS to style my web page and make it look nice.
    - It controls the colors, fonts, spacing, and layout of the elements.

3. **JavaScript**:
    - I use JavaScript to add functionality to my web page.
    - It handles form submissions, processes user input, and updates the results dynamically.

4. **Go (Golang)**:
    - As I have not used go before, I used my limited knowledge of python to deal with my logic and calculation 
    - section and then tried to retrofit all of my code into GoLang.
    - I uses Go for the backend server that processes the data.
    - Go handled my HTTP requests, processesed the data, and then sent back the results.

## Approach

1. **Setting Up the Project**:
    - I created a directory structure to organize my files.
    - I added my HTML, CSS, JavaScript, and Go files.

2. **Creating the HTML Structure**:
    - I added a form for entering pack sizes.
    - I added another form for entering the number of items.
    - I included a div to display the pack size options and another div for the consolidated packs.
    - I placed a logo image above the main heading.

3. **Styling with CSS**:
    - I set the default font to "Bebas Neue" from Google Fonts (As this was apparently Gymsharks main font family(according to a quick google)
    - I center-aligned text and added padding, margins, and borders to make the page look neat and tidy
    - I made the logo responsive so it adjusts its size based on the screen size, also added packages behind the Gymshark logo 😊

4. **Adding JavaScript Functionality**:
    - I handled form submissions to process user input without reloading the page.
    - I displayed the pack size options after the user submits the pack sizes, so the user would always be aware of the
    - selected pack sizes, this allows the user to go back and add more packs if submit incorrect.
    - I calculated the consolidated packs and displayed the results, excluding any packs with a value of zero (as we dont
    - really need to see these)

5. **Creating the Backend with Go**:
    - I set up a Go server to handle HTTP requests.
    - I wrote functions to process the data and calculate the consolidated packs.
    - I sent the results back to the frontend in JSON format.