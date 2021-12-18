package adventofcode.year2021;

import java.util.ArrayList;
import java.util.Stack;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;


import static org.junit.Assert.assertEquals;

public class Day18 extends BaseTest {

	@Test public void runTest() throws Exception {
		NumberSlow reduce = new NumberSlow("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]");
		reduce.reduce();
		assertEquals("[[3,[2,[8,0]]],[9,[5,[7,0]]]]", reduce.toString());
		final NumberSlow reduce0 = new NumberSlow("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]");
		reduce0.reduce();
		assertEquals("[[3,[2,[8,0]]],[9,[5,[7,0]]]]", reduce0.toString());
		reduce = new NumberSlow("[[[[[9,8],1],2],3],4]");
		reduce.reduce();
		assertEquals("[[[[0,9],2],3],4]", reduce.toString());
		NumberSlow reduce1 = new NumberSlow("[7,[6,[5,[4,[3,2]]]]]");
		reduce1.reduce();
		assertEquals("[7,[6,[5,[7,0]]]]", reduce1.toString());
		final NumberSlow reduce2 = new NumberSlow("[[6,[5,[4,[3,2]]]],1]");
		reduce2.reduce();
		assertEquals("[[6,[5,[7,0]]],3]", reduce2.toString());
		NumberSlow ns1 = new NumberSlow("[[[[4,3],4],4],[7,[[8,4],9]]]");
		NumberSlow ns2 = new NumberSlow("[1,1]");
		NumberSlow result = NumberSlow.addNumberSlow(ns1, ns2);
		assertEquals("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", result.toString());
		ns1 = new NumberSlow("[[[[1,1],[2,2]],[3,3]],[4,4]]");
		ns2 = new NumberSlow("[5,5]");
		result = NumberSlow.addNumberSlow(ns1, ns2);
		assertEquals("[[[[3,0],[5,3]],[4,4]],[5,5]]", result.toString());
		reduce = new NumberSlow("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]");
		reduce.reduce();
		assertEquals("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", reduce.toString());
		reduce = new NumberSlow("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]");
		reduce.reduce();
		assertEquals("[[3,[2,[8,0]]],[9,[5,[7,0]]]]", reduce.toString());
		reduce1 = new NumberSlow("[7,[6,[5,[4,[3,2]]]]]");
		reduce1.reduce();
		assertEquals("[7,[6,[5,[7,0]]]]", reduce1.toString());
		ns1 = new NumberSlow("[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]");
		ns2 = new NumberSlow("[2,9]");
		result = NumberSlow.addNumberSlow(ns1, ns2);
		assertEquals("[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]", result.toString());
		ns1 = new NumberSlow("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]");
		ns2 = new NumberSlow("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]");
		result = NumberSlow.addNumberSlow(ns1, ns2);
		assertEquals("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]", result.toString());
		reduce = new NumberSlow("[[[[6,6],[6,7]],[[8,8],[6,[0,7]]]],[7,9]]");
		reduce.reduce();
		assertEquals("[[[[6,6],[6,7]],[[8,8],[6,0]]],[[7,7],9]]", reduce.toString());
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day18/input.txt");
		NumberSlow summ = new NumberSlow(data.get(0));
		for (int i = 1; i < data.size(); i++) {
			final NumberSlow current = new NumberSlow(data.get(i));
			summ = NumberSlow.addNumberSlow(summ, current);
		}
		System.out.println(summ.countMagnitude());
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day18/input.txt");
		int max = Integer.MIN_VALUE;
		for (int i = 0; i < data.size(); i++) {
			for (int j = 0; j < data.size(); j++) {
				if (i != j) {
					final NumberSlow summ = NumberSlow.addNumberSlow(new NumberSlow(data.get(i)), new NumberSlow(data.get(j)));
					max = Math.max(max, summ.countMagnitude());
				}
			}

		}
		System.out.println(max);
	}

	@Data static class NumberSlow {

		NumberSlow parent;
		NumberSlow left;
		NumberSlow right;
		Integer simpleValue;

		public NumberSlow() {}

		public NumberSlow(final String input, final NumberSlow parent) {
			this(input);
			this.parent = parent;
		}

		public NumberSlow(final String input) {
			if (!input.startsWith("[")) {
				this.simpleValue = Integer.parseInt(input);
			} else {
				String rightSubstring = getRightString(input.substring(1));
				if (rightSubstring.startsWith(",")) {
					rightSubstring = rightSubstring.substring(1);
				}
				final String leftSubstring = input.substring(1, input.length() - rightSubstring.length() - 2);
				this.left = new NumberSlow(leftSubstring, this);
				this.right = new NumberSlow(rightSubstring, this);
			}
		}

		public static NumberSlow addNumberSlow(final NumberSlow number1, final NumberSlow number2) {
			final NumberSlow result = new NumberSlow("[" + number1.toString() + "," + number2.toString() + "]");
			result.reduce();
			return result;
		}

		public int countMagnitude() {
			if (isSimple()) {
				return this.simpleValue;
			}
			return 3 * this.left.countMagnitude() + 2 * this.right.countMagnitude();
		}

		public void addToSimpleValue(final int toAdd) {
			this.simpleValue = this.simpleValue + toAdd;
		}

		public void reduce() {
			while (explodeFirst() || splitFirst()) {
			}
		}

		public int countLevel() {
			int level = 0;
			NumberSlow parent = this.parent;
			while (parent != null) {
				level++;
				parent = parent.parent;
			}
			return level;
		}

		public boolean splitFirst() {
			boolean splitted = false;
			if (this.isSimple() && this.simpleValue >= 10) {
				final int leftValue = this.simpleValue / 2;
				final int rightValue = this.simpleValue - leftValue;
				this.left = new NumberSlow();
				this.right = new NumberSlow();
				this.left.parent = this;
				this.left.simpleValue = leftValue;
				this.right.parent = this;
				this.right.simpleValue = rightValue;
				splitted = true;
			}
			if (this.left != null && !splitted) {
				splitted = this.left.splitFirst();
			}
			if (this.right != null && !splitted) {
				splitted = this.right.splitFirst();
			}
			return splitted;
		}

		public boolean explodeFirst() {
			boolean exploded = false;
			if (countLevel() >= 4) {
				if (this.left != null && this.left.isSimple() && this.right != null && this.right.isSimple()) {
					final int toAddLeft = this.left.simpleValue;
					final int toAddRight = this.right.simpleValue;
					this.addToFirstSimpleAny(toAddLeft, false);
					this.addToFirstSimpleAny(toAddRight, true);
					this.left = null;
					this.right = null;
					this.simpleValue = 0;
					exploded = true;
				}
			}
			if (this.left != null && !exploded) {
				exploded = this.left.explodeFirst();
			}
			if (this.right != null && !exploded) {
				exploded = this.right.explodeFirst();
			}
			return exploded;
		}

		private boolean addToFirstSimpleAny(final int toAdd, final boolean isRight) {
			final NumberSlow parent = this.parent;
			boolean added = false;
			if (parent != null) {
				final NumberSlow parentAny = isRight ? parent.right : parent.left;
				if (parentAny.isSimple()) {
					parentAny.addToSimpleValue(toAdd);
					added = true;
				} else {
					NumberSlow nodeToUpdate = null;
					if (parentAny != this) {
						nodeToUpdate = isRight ? parentAny.left : parentAny.right;
					} else {
						NumberSlow parentAnyParent = parentAny.parent;
						NumberSlow cur;
						boolean found = false;
						while (parentAnyParent != null && !found) {
							cur = parentAnyParent;
							parentAnyParent = parentAnyParent.parent;
							if (parentAnyParent != null && (isRight ? parentAnyParent.right != cur : parentAnyParent.left != cur)) {
								nodeToUpdate = isRight ? parentAnyParent.right : parentAnyParent.left;
								found = true;
							}

						}
					}
					if (nodeToUpdate != null) {
						while (nodeToUpdate != null && !added) {
							if (nodeToUpdate.isSimple()) {
								nodeToUpdate.addToSimpleValue(toAdd);
								added = true;
							}
							nodeToUpdate = isRight ? nodeToUpdate.left : nodeToUpdate.right;
						}
					}
				}
			}
			return added;
		}

		private String getRightString(final String s) {
			final Stack<Character> stack = new Stack<>();
			int current = 0;
			for (final char c : s.toCharArray()) {
				current++;
				if (c == '[') {
					stack.push(']');
				} else if (c == ']') {
					stack.pop();
				} else if (stack.isEmpty()) {
					return s.substring(current, s.length() - 1);
				}
			}
			return StringUtils.EMPTY;
		}

		boolean isSimple() {
			return this.left == null && this.right == null;
		}

		@Override public String toString() {
			if (this.isSimple()) {
				return "" + this.simpleValue;
			} else {
				return "[" + this.left + "," + this.right + "]";
			}
		}
	}

}
