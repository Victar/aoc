package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day24 extends BaseTest {

	public static final int DAY = 24;

	char[][] map;
	Map<Integer, List<Blizzard>> mapBlizzardsList = new HashMap<>();

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runBoth() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		map = new char[data.size()][data.get(0).length()];
		Point start = new Point(1, 0);
		Point goal = new Point(data.get(0).length() - 2, data.size() - 1);
		List<Blizzard> blizzardsList = new ArrayList<>();
		for (int y = 0; y < data.size(); y++) {
			for (int x = 0; x < data.get(0).length(); x++) {
				char current = data.get(y).charAt(x);
				map[y][x] = current;
				if (current == '>' || current == '<' || current == '^' || current == 'v') {
					blizzardsList.add(new Blizzard(new Point(x, y), current));
					map[y][x] = '.';
				}
			}
		}
		mapBlizzardsList.put(0, blizzardsList);

		Point startOne = new Point(start.x, start.y, null, 0);
		int steps1 = findPath(startOne, goal);
		System.out.println("Silver:" + steps1);

		Point startTwo = new Point(goal.x, goal.y, null, steps1);
		int steps2 = findPath(startTwo, start);

		Point startThree = new Point(start.x, start.y, null, steps2);
		int steps3 = findPath(startThree, goal);
		System.out.println("Gold: " + steps3);

	}

	public int findPath(Point start, Point end) {
		Deque<Point> queue = new ArrayDeque<>();
		Set<String> visited = new HashSet<>();
		int steps = start.step;
		queue.add(start);
		visited.add(start.getState());
		while (!queue.isEmpty()) {
			int size = queue.size();
			for (int i = 0; i < size; i++) {
				Point curr = queue.poll();
				if (curr.x == end.x && curr.y == end.y) {
					return steps;
				}
				Point[] neighbors = new Point[] {
						new Point(curr.x, curr.y, curr, steps + 1), // wait
						new Point(curr.x, curr.y - 1, curr, steps + 1), // up
						new Point(curr.x, curr.y + 1, curr, steps + 1), // down
						new Point(curr.x - 1, curr.y, curr, steps + 1), // left
						new Point(curr.x + 1, curr.y, curr, steps + 1), // right
				};
				for (Point next : neighbors) {
					if (next.x >= 0 && next.x < map[0].length && next.y >= 0 && next.y < map.length && isValidAtStep(next)
							&& !visited.contains(next.getState())) {
						queue.add(next);
						visited.add(next.getState());
					}
				}
			}
			steps++;
		}
		return -1;
	}

	public void print(Point point) {
		System.out.println("   " + point + "    ");
		for (int y = 0; y < map.length; y++) {
			for (int x = 0; x < map[0].length; x++) {
				Point p = new Point(x, y, null, point.step);
				if (p.x == point.x && p.y == point.y) {
					System.out.print("E");
				} else {
					Blizzard blizzard = blizzardsAtStep(p);
					if (blizzard != null) {
						System.out.print(blizzard.dir);
					} else {
						System.out.print(map[y][x]);
					}
				}
			}
			System.out.println();
		}
	}

	public List<Blizzard> getUpdatedBlizzards(int step) {
		if (mapBlizzardsList.containsKey(step)) {
			return mapBlizzardsList.get(step);
		} else {

			List<Blizzard> blizzardsListPrevious = getUpdatedBlizzards(step - 1);

			List<Blizzard> blizzardsListCopy = new ArrayList<>();
			for (Blizzard b : blizzardsListPrevious) {
				blizzardsListCopy.add(b.clone());
			}

			for (Blizzard b : blizzardsListCopy) {
				switch (b.dir) {
					case '^':
						b.pos.y--;
						break;
					case 'v':
						b.pos.y++;
						break;
					case '<':
						b.pos.x--;
						break;
					case '>':
						b.pos.x++;
						break;
				}
				if (b.pos.x < 1) b.pos.x = map[0].length - 2;
				if (b.pos.x >= map[0].length - 1) b.pos.x = 1;
				if (b.pos.y < 1) b.pos.y = map.length - 2;
				if (b.pos.y >= map.length - 1) b.pos.y = 1;
			}
			mapBlizzardsList.put(step, blizzardsListCopy);
			return blizzardsListCopy;
		}
	}

	public boolean isValidAtStep(Point p) {
		return map[p.y][p.x] != '#' && !isOccupiedAtStep(p);
	}

	public boolean isOccupiedAtStep(Point p) {
		List<Blizzard> blizzardsListNextMinute = getUpdatedBlizzards(p.step);
		for (Blizzard b : blizzardsListNextMinute) {
			if (b.pos.x == p.x && b.pos.y == p.y) {
				return true;
			}
		}
		return false;
	}

	public Blizzard blizzardsAtStep(Point p) {
		List<Blizzard> blizzardsListNextMinute = getUpdatedBlizzards(p.step);
		for (Blizzard b : blizzardsListNextMinute) {
			if (b.pos.x == p.x && b.pos.y == p.y) {
				return b;
			}
		}
		return null;
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (final String input : data) {
			System.out.println(input);
		}
	}

	@Data static class Blizzard {

		Point pos;
		char dir;

		public Blizzard(Point pos, char dir) {
			this.pos = pos;
			this.dir = dir;
		}

		public Blizzard clone() {
			return new Blizzard(new Point(this.pos.x, this.pos.y), dir);
		}
	}

	@Data static class Point {

		int x, y;
		Point prev;
		int step;

		public Point(int x, int y) {
			this.x = x;
			this.y = y;
		}

		public Point(int x, int y, Point prev, int step) {
			this.x = x;
			this.y = y;
			this.prev = prev;
			this.step = step;
		}

		@Override public boolean equals(Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;
			Point point = (Point) o;
			return x == point.x && y == point.y;
		}

		@Override public int hashCode() {
			return Objects.hash(x, y);
		}

		@Override public String toString() {
			return "{" + "x=" + x + ", y=" + y + ", step=" + step + ", prev=" + prev + '}';
		}

		public String getState() {
			return x + ":" + y + ":" + step;
		}
	}

}
