using System.Text.Json.Serialization;

namespace ChuckJokes.WebApp.Models
{
    public class Joke
    {
        public int Id { get; set; }
        [JsonPropertyName("joke")]
        public string Content { get; set; }
        public bool Explicit { get; set; }
    }
}