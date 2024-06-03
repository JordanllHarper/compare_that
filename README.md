# The Idea
Be able to enter in a product of some kind. It has a title, price, rating, and some notes the user can make. 

Schema: Item {

    name string

    price int (lowest form of currency)

    rating int (user defined rating scale) could be represented as '*' in gui

    notes string

}

A user can type in an ID of an item (which is randomly assigned to it) and get the info they captured on that item.

A user can also compare 2 or more items with `compare x_id y_id`, where x and y are different products, in a nice side by side view which shows 
price differences. 


----
