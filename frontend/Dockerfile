# ---------- Build stage ----------
FROM node:20-alpine AS build
WORKDIR /app

# Install deps
COPY package*.json ./
RUN npm ci

# Build
COPY . .
RUN npm run build

# ---------- Run stage ----------
FROM nginx:1.25-alpine
# Copy static assets
COPY --from=build /app/dist /usr/share/nginx/html
# Nginx config (SPA + /api proxy)
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]