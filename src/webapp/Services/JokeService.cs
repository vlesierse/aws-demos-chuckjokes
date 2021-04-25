using System.Collections.Generic;
using System.Net.Http;
using System.Text.Json;
using System.Threading.Tasks;
using ChuckJokes.WebApp.Models;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;

namespace ChuckJokes.WebApp.Services
{
    public class JokeService : IJokeService
    {
        private readonly HttpClient _httpClient;
        private readonly IOptions<AppSettings> _settings;
        private readonly ILogger<JokeService> _logger;

        public JokeService(HttpClient httpClient, IOptions<AppSettings> settings, ILogger<JokeService> logger)
        {
            _httpClient = httpClient;
            _settings = settings;
            _logger = logger;
        }
        
        public async Task<Joke> GetRandomJoke()
        {
            var uri = $"{_settings.Value.JokeGeneratorUrl}/jokes/random";
            try
            {
                var responseStream = await _httpClient.GetStreamAsync(uri);
                var joke = await JsonSerializer.DeserializeAsync<Joke>(responseStream,
                    new JsonSerializerOptions {
                        DictionaryKeyPolicy = JsonNamingPolicy.CamelCase,
                        PropertyNamingPolicy = JsonNamingPolicy.CamelCase
                    });
                return joke;
            }
            catch (HttpRequestException e)
            {
                _logger.LogError($"HTTP Request Failed: {e.Message}");
            }
            return null;
        }
    }
}