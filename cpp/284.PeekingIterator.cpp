// Below is the interface for Iterator, which is already defined for you.
// **DO NOT** modify the interface for Iterator.

class Iterator {
    struct Data;
	Data* data;
public:
	Iterator(const vector<int>& nums);
	Iterator(const Iterator& iter);
	virtual ~Iterator();
	// Returns the next element in the iteration.
	int next();
	// Returns true if the iteration has more elements.
	bool hasNext() const;
};


class PeekingIterator : public Iterator {
private:
    int savedNum;
    bool isSaved;
public:
	PeekingIterator(const vector<int>& nums) : Iterator(nums) {
	    // Initialize any member here.
	    // **DO NOT** save a copy of nums and manipulate it directly.
	    // You should only use the Iterator interface methods.
	    isSaved = false;
	}

    // Returns the next element in the iteration without advancing the iterator.
	int peek() {
        if (isSaved) {
            return savedNum;
        }
        savedNum = Iterator::next();
        isSaved = true;
        return savedNum;
	}

	// hasNext() and next() should behave the same as in the Iterator interface.
	// Override them if needed.
	int next() {
	    if (isSaved) {
            isSaved = false;
            return savedNum;
        }
        return Iterator::next();
	}

	bool hasNext() const {
	    return isSaved ? true : Iterator::hasNext();
	}
};

