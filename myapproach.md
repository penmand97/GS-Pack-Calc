Sure, here is your document in UK English and with a less corporate tone:

# Pack Consolidation Web App

Welcome to my submission for the Gymshark Pack Consolidation Assessment!

This guide is here to help you understand my thought process, the methods I used, and how I approached this project. I’ve explained everything in simple terms, so it's easy to follow along.

## Thought Process

1. **Understanding the Problem**:
    - I needed to create a tool to help pack items efficiently using different pack sizes.
    - The tool also needed to show the pack sizes we have and the final consolidated packs without showing packs with a value of zero.

2. **Breaking Down the Task**:
    - First, I needed to get input from the user for the pack sizes and the number of items.
    - Then, I processed this input to work out the best way to pack the items.
    - Finally, I displayed the results on the web page.

3. **User Interface Design**:
    - My web page had two forms: one for entering pack sizes and another for entering the number of items.
    - I showed the pack size options and the consolidated packs as results on the same page.
    - I also included a responsive logo at the top of the page.

## Methods Used

1. **HTML (HyperText Markup Language)**:
    - I used HTML to create the structure of the web page.
    - It defined elements like headings, forms, input fields, and buttons.

2. **CSS (Cascading Style Sheets)**:
    - I used CSS to style the web page and make it look nice.
    - It controlled the colours, fonts, spacing, and layout of the elements.

3. **JavaScript**:
    - I used JavaScript to add functionality to the web page.
    - It handled form submissions, processed user input, and updated the results dynamically.

4. **Go (Golang)**:
    - Since I hadn’t used Go before, I leaned on my limited Python knowledge for the logic and calculations.
    - I then retrofitted all my code into Go.
    - I used Go for the backend server to process the data.
    - Go handled my HTTP requests, processed the data, and sent back the results.

## Approach

1. **Setting Up the Project**:
    - I created a directory structure to organise my files.
    - I added my HTML, CSS, JavaScript, and Go files.

2. **Creating the HTML Structure**:
    - I added a form for entering pack sizes.
    - I added another form for entering the number of items.
    - I included a div to display the pack size options and another div for the consolidated packs.
    - I placed a logo image above the main heading.

3. **Styling with CSS**:
    - I set the default font to "Bebas Neue" from Google Fonts (apparently Gymshark's main font family according to a quick Google search).
    - I centre-aligned text and added padding, margins, and borders to make the page look neat and tidy.
    - I made the logo responsive so it adjusted its size based on the screen size and added some packages behind the Gymshark logo 😊.

4. **Adding JavaScript Functionality**:
    - I handled form submissions to process user input without reloading the page.
    - I displayed the pack size options after the user submitted the pack sizes, so they could always see the selected pack sizes and go back to add more if they submitted incorrectly.
    - I calculated the consolidated packs and displayed the results, excluding any packs with a value of zero (since we didn’t need to see these).

5. **Creating the Backend with Go**:
    - I set up a Go server to handle HTTP requests.
    - I wrote functions to process the data and calculate the consolidated packs.
    - I sent the results back to the frontend in JSON format.
