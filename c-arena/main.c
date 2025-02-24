#include <stdlib.h>
#include <stdio.h>

struct arena {
	char *data;
	size_t capacity;
	size_t offset;
};

#define make(X) _Generic(X, \
	Arena: make_Arena \
)()

#define drop(X) _Generic((X), \
	Arena *: drop_Arena \
)(X)

#define new(X) _Generic(size_t, \
		size_t: arena_alloc(arena, (X)) \
)

typedef struct arena A;
typedef struct arena Arena;

// 64 GiB; virtual mem, does not actually allocate that much
constexpr long long ArenaSize = 68719476736;

static void arena_init(struct arena* a) {
	a->data = malloc(ArenaSize);
	a->capacity = ArenaSize;
	a->offset = 0;
}

Arena *make_Arena() {
	Arena *a = malloc(sizeof(Arena));
	arena_init(a);
	return (Arena*)a;
}

void drop_Arena(Arena *a) {
	free(a->data);
}

void *arena_alloc(struct arena* arena, size_t size) {
	if (arena->offset + size > arena->capacity)
		return nullptr;

	void *ptr = arena->data + arena->offset;
    arena->offset += size;
    return ptr;
}

void bar(A *arena, int N) {
	char **strings = new(sizeof(char*));

	for (int i=0; i < N; i++) {
		char *ptr = new(sizeof(char));
		ptr[0] = 'f';
		*strings++ = ptr;
	}
}

int main(void) {
	Arena *arena = make(Arena);
	printf("capacity: %zu, free: %zu\n",
			arena->capacity,
			arena->capacity - arena->offset);

	bar(arena, 100);

	printf("capacity: %zu, free: %zu\n",
			arena->capacity,
			arena->capacity - arena->offset);

	while(true) {}

	drop(arena);
}
