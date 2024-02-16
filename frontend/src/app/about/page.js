import Milestone from "@/components/MileStone";
import Image from "next/image";

export default function About() {
    return (
        <div className="bg-gray-200 w-full max-w-container mx-auto p-4 mt-2 rounded">
            <div className="flex items-center">
                <div className="relative h-48 w-48 flex-logo">
                    <Image
                        src="/profile.jpg"
                        className="rounded-full object-cover object-center border-gray-100"
                        fill
                        alt="Profile"
                    />
                </div>
                <div className="pl-4">
                    <h1 className="text-4xl font-bold">T1 Gumayusi</h1>
                    <div className="text-justify mt-2">
                        Aenean nibh lacus, dignissim sed nisi eget, posuere luctus lorem. Vivamus sagittis metus eget auctor volutpat. In ut leo sollicitudin tellus lobortis congue. Aliquam vel nisi tortor. Suspendisse ac finibus odio. Nunc sagittis enim urna. Donec augue mauris, varius quis convallis et, aliquam vel velit. Suspendisse nec lorem sit amet diam laoreet pellentesque et non lectus. Sed lacus mauris, porttitor nec molestie at, tincidunt mattis sapien. Quisque quam tortor, blandit vitae mauris vitae, pretium consequat velit. Sed nisi lacus, volutpat nec gravida eu, tristique ac sem. Proin elementum turpis et nunc molestie placerat. Phasellus a nisl eget nibh porta ultrices.
                    </div>
                </div>
            </div>
            <div className="h-1 my-4 w-full bg-gray-50"></div>
            <div>
                <h1 className="text-3xl font-semibold mb-4">Brief history</h1>
                <Milestone year={2019} title="Bot laner at T1 Current Roster" content="Nullam vitae mi urna. Cras augue lorem, aliquet sit amet scelerisque quis, elementum eget dui. Mauris eu lacus et massa tempor interdum. Ut pretium tristique blandit. Nam in nisl condimentum, lacinia tortor vulputate, tristique sapien. Pellentesque porta tellus a dignissim sagittis. In aliquam volutpat lacus vel ornare." />
                <Milestone year={2018} title="Bot laner at Sk Telecom T1" content="Nullam vitae mi urna. Cras augue lorem, aliquet sit amet scelerisque quis, elementum eget dui. Mauris eu lacus et massa tempor interdum. Ut pretium tristique blandit. Nam in nisl condimentum, lacinia tortor vulputate, tristique sapien. Pellentesque porta tellus a dignissim sagittis. In aliquam volutpat lacus vel ornare." />
                <Milestone year={2018} title="Bot laner at Seoul Current Roster" content="Nullam vitae mi urna. Cras augue lorem, aliquet sit amet scelerisque quis, elementum eget dui. Mauris eu lacus et massa tempor interdum. Ut pretium tristique blandit. Nam in nisl condimentum, lacinia tortor vulputate, tristique sapien. Pellentesque porta tellus a dignissim sagittis. In aliquam volutpat lacus vel ornare." />
            </div>
        </div>
    );
}