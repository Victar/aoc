package adventofcode.year2021;

import java.util.*;
import java.util.stream.Collectors;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day24 extends BaseTest {

	private static final int INSTRUCTIONS_COUNT = 14;

	Map<Integer, Integer> mapping = new HashMap<>();

	@Test public void runSilver() throws Exception {
		runAny(true);
	}

	@Test public void runGold() throws Exception {
		runAny(false);
	}

	// Non-trivial task was simplified
	// more details https://www.reddit.com/r/adventofcode/comments/rnejv5/2021_day_24_solutions/
	public void runAny(final boolean max) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day24/input.txt");
		final List<Integer> check = new ArrayList<>();
		final List<Integer> offset = new ArrayList<>();
		final int instructionSize = data.size() / INSTRUCTIONS_COUNT;
		for (int i = 0; i < INSTRUCTIONS_COUNT; i++) {
			check.add(Integer.parseInt(data.get(i * instructionSize + 5).substring(6)));
			offset.add(Integer.parseInt(data.get(i * instructionSize + 15).substring(6)));
		}
		final List<Pair> pairs = new ArrayList<>();
		final Stack<Integer> stack = new Stack<>();
		for (int i = 0; i < INSTRUCTIONS_COUNT; i++) {
			if (check.get(i) < 0) {
				final Integer position2 = stack.pop();
				pairs.add(new Pair(i, position2, offset.get(position2) + check.get(i)));
			} else if (offset.get(i) > 0) {
				stack.push(i);
			}
		}
		System.out.println(solve(pairs, max));
	}

	String solve(final List<Pair> pairs, final boolean max) {
		final Integer[] positions = new Integer[INSTRUCTIONS_COUNT];
		for (final Pair pair : pairs) {
			pair.updatePosition(positions, max);
		}
		return Arrays.stream(positions).map(String::valueOf).collect(Collectors.joining());
	}

	@Data class Pair {

		int position1;
		int position2;
		int offset;

		public Pair(final int position1, final int position2, final int offset) {
			this.position1 = position1;
			this.position2 = position2;
			this.offset = offset;
		}

		public void updatePosition(final Integer[] position, final boolean max) {
			if (max) {
				if (this.offset > 0) {
					position[this.position1] = 9;
					position[this.position2] = 9 - this.offset;
				} else {
					position[this.position1] = 9 + this.offset;
					position[this.position2] = 9;
				}
			} else {
				if (this.offset > 0) {
					position[this.position1] = 1 + this.offset;
					position[this.position2] = 1;
				} else {
					position[this.position1] = 1;
					position[this.position2] = 1 - this.offset;
				}
			}
		}

		@Override public String toString() {
			return "[" + this.position1 + "] = [" + this.position2 + "]  " + this.offset;
		}
	}

}
