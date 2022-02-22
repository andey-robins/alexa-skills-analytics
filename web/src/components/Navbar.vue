<template>
<nav class="bg-white shadow-lg">
    <div class="max-w-6xl mx-auto px-4">
        <div class="flex justify-between">
            <div class="flex space-x-7">
                <div>
                    <a href="#" class="navbar-a flex items-center py-4 px-2 float-left">
                        <img src="logo.png" alt="Logo" class="h-8 w-8 mr-2" />
                        <span class="font-semibold text-lg">Department of Education</span>
                    </a>
				</div>
                <div id="nav" class="hidden md:flex items-center space-x-1 float-right text-gray-500">
                    <a class="navbar-a" href="#skills"><p>Skills</p></a>
                    <a class="navbar-a" href="#about"><p>About</p></a>
                </div>
            </div>
        </div>
    </div>
    
</nav>
</template>

<style>
.navbar-a {
    @apply py-4 px-2 text-amber-900 border-b-4 border-amber-900 font-semibold
}
</style>