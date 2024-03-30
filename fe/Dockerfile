# Use Node.js LTS version as the base image
FROM node:lts-alpine as builder

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application
COPY . .

# Build the Angular app for production
RUN npm run build --prod

# Create a new image with nginx
FROM node:lts-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the built Angular app from the builder stage
COPY --from=builder /app/dist/ditto .

# Install a simple HTTP server
RUN npm install -g http-server

# Expose port 8080
EXPOSE 8080

# Command to run the Angular app
CMD ["http-server", "-p", "8080"]
