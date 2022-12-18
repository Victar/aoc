package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day18 extends BaseTest {

	public static final int DAY = 18;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		List<Dot> dots = new ArrayList<>();
		for (final String input : data) {
			dots.add(new Dot(input));
		}
		int result = 0;
		for (Dot dot : dots) {
			result += dot.getBordersWithoutNeighbors(dots);
		}
		System.out.println(result);

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		List<Dot> dots = new ArrayList<>();
		Dot min = new Dot(Integer.MAX_VALUE, Integer.MAX_VALUE, Integer.MAX_VALUE);
		Dot max = new Dot(Integer.MIN_VALUE, Integer.MIN_VALUE, Integer.MIN_VALUE);

		for (final String input : data) {
			Dot current = new Dot(input);
			min = new Dot(Math.min(min.x, current.x), Math.min(min.y, current.y), Math.min(min.z, current.z));
			max = new Dot(Math.max(max.x, current.x), Math.max(max.y, current.y), Math.max(max.z, current.z));
			dots.add(current);
		}
		int result = 0;
		for (Dot dot : dots) {
			List<Dot> neighbors = dot.getNeighbors();
			for (Dot neighbor : neighbors){
				if (canMoveOut(neighbor, dots, min, max)){
					result ++;
				}
			}
		}
		System.out.println(result);
	}

	boolean canMoveOut(Dot dotToCheck, List<Dot> dots, Dot min, Dot max) {
		final Stack<Dot> stack = new Stack<>();
		stack.push(dotToCheck);

		Set<Dot> visited = new HashSet<>();
		visited.add(dotToCheck);

		while (!stack.isEmpty()) {
			Dot current = stack.pop();
			if (dots.contains(current)) {
				continue;
			}

			if (current.isOutside(min, max)) {
				return true;
			}
			List<Dot> neighbors = current.getNeighbors();
			for (Dot neighbor : neighbors) {
				if (!visited.contains(neighbor)) {
					visited.add(neighbor);
					stack.push(neighbor);
				}
			}
		}
		return false;
	}

	@Data static class Dot {

		int x;
		int y;
		int z;

		public Dot(String input) {
			final String[] args = input.split(",");
			x = Integer.parseInt(args[0]);
			y = Integer.parseInt(args[1]);
			z = Integer.parseInt(args[2]);
		}

		public Dot(int x, int y, int z) {
			this.x = x;
			this.y = y;
			this.z = z;
		}

		boolean isOutside(Dot min, Dot max) {
			return this.x < min.x || this.x > max.x //
					|| this.y < min.y || this.y > max.y //
					|| this.z < min.z || this.z > max.z; //

		}

		public List<Dot> getNeighbors() {
			List<Dot> neighbors = new ArrayList<>();
			for (int x = -1; x <= 1; x++) {
				for (int y = -1; y <= 1; y++) {
					for (int z = -1; z <= 1; z++) {
						if ((Math.abs(x) + Math.abs(y) + Math.abs(z)) == 1) {
							neighbors.add(new Dot(this.x + x, this.y + y, this.z + z));
						}
					}
				}
			}
			return neighbors;
		}

		public int getBordersWithoutNeighbors(List<Dot> dots) {
			List<Dot> neighbors = getNeighbors();
			int count = 0;
			for (Dot neighbor : neighbors) {
				if (!dots.contains(neighbor)) {
					count++;
				}
			}
			return count;
		}
	}

}
