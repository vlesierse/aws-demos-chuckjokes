FROM mcr.microsoft.com/dotnet/aspnet:5.0 AS base
WORKDIR /app
EXPOSE 80

FROM --platform=$BUILDPLATFORM mcr.microsoft.com/dotnet/sdk:5.0 AS build
WORKDIR /src

COPY "ChuckJokes.WebApp.csproj" "ChuckJokes.WebApp.csproj"

RUN dotnet restore "ChuckJokes.WebApp.csproj"

COPY . .
RUN dotnet publish --no-restore -c Release -o /app

FROM base AS final
WORKDIR /app
COPY --from=build /app .
ENTRYPOINT ["dotnet", "ChuckJokes.WebApp.dll"]
