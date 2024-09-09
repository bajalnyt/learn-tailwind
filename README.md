
## Running without docker:

In 2 separate terminals, run:
```
task run
task tailwind-watch 
```

## Initial setup for tailwind
1. Install Tailwind CSS
```
npm init -y
npm install -D tailwindcss
npx tailwindcss init
```
2. Configure your template paths
Create `tailwind.config.js`

3. Create src/input.css

4. Start tailwind cli build process
`npx tailwindcss -i ./src/input.css -o ./src/output.css --watch`
