<template>
<nav className="bg-white shadow-lg">
    <div className="max-w-6xl mx-auto px-4">
        <div className="flex justify-between">
            <div className="flex space-x-7">
                <div>
                    <a href="#" className="navbar-a flex items-center py-4 px-2 float-left">
                        <span className="font-semibold text-lg">Department of Education</span>
                    </a>
				</div>
                <div id="nav" className="hidden md:flex items-center space-x-1 float-right text-gray-500">
                    <a className="navbar-a" href="#skills"><p>Skills</p></a>
                    <a className="navbar-a" href="#about"><p>About</p></a>
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