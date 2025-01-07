
# 🚀 **Supplement and Hydration Calculator**

A modern **web-based Supplement and Hydration Calculator** designed to empower users to calculate their daily **protein intake**, **creatine dosage**, **whey protein**, and **water requirements** based on their **weight** and **activity level**. Supporting **multiple languages** (English, Portuguese, and Spanish), it ensures a seamless and inclusive user experience.

---

## 📚 **Description**

This project helps users optimize their health and fitness routines by providing accurate nutritional and hydration recommendations. With a dynamic interface and language localization, users across different regions can benefit from personalized results.

---

## 🚀 **Key Features**

### ✅ **Dynamic Supplement Calculator**
- Calculate daily requirements for:
   - **Protein intake**
   - **Creatine dosage**
   - **Whey protein**
   - **Water intake**
- Inputs based on **user weight** and **activity level**.

### ✅ **Multi-Language Support**
- Available in:
   - **English (en-us)**
   - **Portuguese (pt-br)**
   - **Spanish (es-es)**
- Language switching through an intuitive **dropdown menu**.

### ✅ **Modern and Responsive Interface**
- Built with **Tailwind CSS** for a clean and modern look.
- Fully **responsive design**, optimized for **mobile** and **desktop**.

### ✅ **SEO Optimized**
- Dynamic `<meta>` tags for **title**, **description**, and **keywords**.
- Proper **Open Graph (OG)** and **Twitter Card** integration for enhanced sharing.

### ✅ **Server-Side Rendering (SSR)**
- **Go Templates** ensure fast and dynamic rendering on the server.
- Backend powered by **Go (Golang)** for high efficiency.

### ✅ **JSON-Based Translations**
- Language content is managed through structured **JSON files**:
   - `en.json`
   - `pt_br.json`
   - `es.json`
- Easy updates and maintenance.

---

## ⚙️ **Tech Stack**

- **Backend:** Go (Golang)
- **Frontend:** Tailwind CSS
- **Templating Engine:** Go Templates
- **Data Storage:** JSON files (for translations)
- **Web Server:** Go HTTP Server

---

## 📂 **Project Structure**

\`\`\`plaintext
├── locales/
│   ├── en.json       # English Translations
│   ├── pt_br.json    # Portuguese Translations
│   ├── es.json       # Spanish Translations
├── static/
│   ├── js/
│   │   ├── scripts.js # Frontend JavaScript
├── templates/
│   ├── head.html     # Meta and Head Information
│   ├── index.html    # Main Page Template
│   ├── selectLang.html # Language Selector Component
├── main.go          # Application Entry Point
├── go.mod           # Go Module Configuration
└── .gitignore       # Ignored Files
\`\`\`

---

## 🛠️ **Getting Started**

Follow the steps below to set up and run the **Supplement and Hydration Calculator** on your local machine.

### 🔹 **1. Clone the Repository**

Start by cloning the project from GitHub and navigating to the project directory:

\`\`\`bash
git clone https://github.com/itsjhonny/go-template-multi-languages.git
cd go-template-multi-languages
\`\`\`

---

### 🔹 **2. Install Dependencies**

Make sure all Go dependencies are properly installed:

\`\`\`bash
go mod tidy
\`\`\`

This command ensures that all required packages are fetched and the `go.mod` file is up to date.

---

### 🔹 **3. Run the Application**

Start the Go HTTP server to serve the application:

\`\`\`bash
go run main.go
\`\`\`

If everything is set up correctly, you will see an output like this:

\`\`\`plaintext
Servidor rodando em http://localhost:8080
\`\`\`

---

### 🔹 **4. Access the Application**

Open your browser and navigate to:

👉 **[http://localhost:8080](http://localhost:8080)**

You can access the calculator in different languages using the following routes:

- 🟢 **English:** `http://localhost:8080/en-us/calc`
- 🟢 **Portuguese (BR):** `http://localhost:8080/pt-br/calc`
- 🟢 **Spanish:** `http://localhost:8080/es-es/calc`

---

### 🔹 **5. Change Language Dynamically**

You can also change the language dynamically using query strings:

- 🟢 **English:** [http://localhost:8080/set-lang?lang=en-us](http://localhost:8080/set-lang?lang=en-us)
- 🟢 **Portuguese (BR):** [http://localhost:8080/set-lang?lang=pt-br](http://localhost:8080/set-lang?lang=pt-br)
- 🟢 **Spanish:** [http://localhost:8080/set-lang?lang=es-es](http://localhost:8080/set-lang?lang=es-es)

---

### 🔹 **6. Stop the Server**

To stop the server, press:

\`\`\`bash
Ctrl + C
\`\`\`

---

## 📚 **Next Steps**

- 🛠️ Explore the different language options.
- 🛠️ Test the calculator with various inputs.
- 🛠️ Dive into the codebase and feel free to contribute!

---

🚀 **Enjoy your journey with the Supplement and Hydration Calculator!** 🌟
