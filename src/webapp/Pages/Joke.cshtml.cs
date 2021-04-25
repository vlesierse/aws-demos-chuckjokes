using System.Threading.Tasks;
using ChuckJokes.WebApp.Services;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace ChuckJokes.WebApp.Pages
{
    public class JokeModel : PageModel
    {
        private readonly IJokeService _jokeService;

        public JokeModel(IJokeService jokeService)
        {
            _jokeService = jokeService;
        }

        public string Joke { get; private set; }
        
        public async Task OnGetAsync()
        {
            var joke = await _jokeService.GetRandomJoke();
            Joke = joke?.Content ?? "Oeps, I can't find a joke. This is embarrassing and not funny at all.";
        }
    }
}